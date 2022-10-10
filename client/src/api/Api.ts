/* eslint-disable */
/* tslint:disable */
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface GormDeletedAt {
  time?: string;

  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface ModelChannel {
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  description?: string;
  id?: number;
  name?: string;
  owner?: number;
  picture?: string;
  private?: boolean;
  updatedAt?: string;
  users?: ModelChannelUser[];
}

export interface ModelChannelUser {
  canMod?: boolean;
  canRead?: boolean;
  canSend?: boolean;
  channelID?: number;
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  id?: number;
  updatedAt?: string;
  user?: ModelPublicUser;
  userID?: number;
}

export interface ModelMessage {
  channelID?: number;
  content?: string;
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  id?: number;
  updatedAt?: string;
  userID?: number;
}

export interface ModelPublicUser {
  email?: string;
  name?: string;
  picture?: string;
  pseudo?: string;
  role?: string;
  surname?: string;
}

export interface ModelUser {
  channels?: ModelChannelUser[];
  createdAt?: string;
  deletedAt?: GormDeletedAt;
  email?: string;
  id?: number;
  myChannels?: ModelChannel[];
  name?: string;
  picture?: string;
  pseudo?: string;
  role?: string;
  surname?: string;
  updatedAt?: string;
}

export interface RouteCheckChanNameBody {
  name?: string;
}

export interface RouteCheckChanNameReturn {
  available?: boolean;
  error?: boolean;
  message?: string;
}

export interface RouteCreateChanBody {
  description?: string;
  name?: string;
  picture?: string;
  private?: boolean;
}

export interface RouteCreateChanReturn {
  error?: boolean;
  message?: string;
  passkey?: string;
  updated?: boolean;
}

export interface RouteDeleteChanReturn {
  deleted?: boolean;
  error?: boolean;
  message?: string;
}

export interface RouteGetOneChanReturn {
  error?: boolean;
  message?: string;
}

export interface RouteGetUserChanReturn {
  memberChan?: ModelChannel[];
  ownChan?: ModelChannel[];
}

export interface RouteJoinChanReturn {
  error?: boolean;
  joined?: boolean;
  message?: string;
}

export interface RouteLoginBody {
  email?: string;
  password?: string;
}

export interface RouteLoginReturn {
  access_token?: string;
  pseudo?: string;
  renew_token?: string;
  role?: string;
}

export interface RouteLogoutBody {
  renew_token?: string;
}

export interface RouteRegisterBody {
  email?: string;
  name?: string;
  password?: string;
  pseudo?: string;
  surname?: string;
}

export interface RouteRegisterReturn {
  error?: boolean;
  message?: string;
  registered?: boolean;
}

export interface RouteRenewChanPasswordReturn {
  error?: boolean;
  message?: string;
  updated?: boolean;
}

export interface RouteRenewTokenBody {
  renew_token?: string;
}

export interface RouteRenewTokenReturn {
  access_token?: string;
  pseudo?: string;
  role?: string;
}

export interface RouteSendMessageContent {
  content?: string;
}

export interface RouteSetPictureReturn {
  message?: string;
  success?: boolean;
}

export interface RouteSetUserBody {
  name?: string;
  surname?: string;
}

export interface RouteSetUserReturn {
  error?: boolean;
  message?: string;
  updated?: boolean;
}

export interface RouteUpdateChannelBody {
  description?: string;
  private?: boolean;
}

export interface RouteUpdateChannelReturn {
  error?: boolean;
  message?: string;
  updated?: boolean;
}

export interface RouteUpdateUserChanRightBody {
  canMod?: boolean;
  canRead?: boolean;
  canSend?: boolean;
  userId?: number;
}

export interface RouteUpdateUserChanRightReturn {
  error?: boolean;
  message?: string;
  updated?: boolean;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<
  FullRequestParams,
  "body" | "method" | "query" | "path"
>;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (
    securityData: SecurityDataType | null
  ) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown>
  extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "/api";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) =>
    fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(
      typeof value === "number" ? value : `${value}`
    )}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter(
      (key) => "undefined" !== typeof query[key]
    );
    return keys
      .map((key) =>
        Array.isArray(query[key])
          ? this.addArrayQueryParam(query, key)
          : this.addQueryParam(query, key)
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string")
        ? JSON.stringify(input)
        : input,
    [ContentType.FormData]: (input: any) =>
      Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
            ? JSON.stringify(property)
            : `${property}`
        );
        return formData;
      }, new FormData()),
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(
    params1: RequestParams,
    params2?: RequestParams
  ): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (
    cancelToken: CancelToken
  ): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(
      `${baseUrl || this.baseUrl || ""}${path}${
        queryString ? `?${queryString}` : ""
      }`,
      {
        ...requestParams,
        headers: {
          ...(type && type !== ContentType.FormData
            ? { "Content-Type": type }
            : {}),
          ...(requestParams.headers || {}),
        },
        signal: cancelToken
          ? this.createAbortSignal(cancelToken)
          : requestParams.signal,
        body:
          typeof body === "undefined" || body === null
            ? null
            : payloadFormatter(body),
      }
    ).then(async (response) => {
      const r = response as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title Bipper Api
 * @version 1.0
 * @baseUrl /api
 * @contact Batleforc <maxleriche.60@gmail.com> (https://weebo.fr)
 *
 * Bipper api
 */
export class Api<
  SecurityDataType extends unknown
> extends HttpClient<SecurityDataType> {
  asset = {
    /**
     * @description Serve static asset
     *
     * @tags Asset
     * @name AssetDetail
     * @summary Serve static asset
     * @request GET:/asset/{fileName}
     */
    assetDetail: (fileName: string, params: RequestParams = {}) =>
      this.request<File, any>({
        path: `/asset/${fileName}`,
        method: "GET",
        ...params,
      }),
  };
  auth = {
    /**
     * @description Login user
     *
     * @tags Auth
     * @name LoginCreate
     * @summary Login user
     * @request POST:/auth/login
     */
    loginCreate: (Request: RouteLoginBody, params: RequestParams = {}) =>
      this.request<RouteLoginReturn, any>({
        path: `/auth/login`,
        method: "POST",
        body: Request,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Logout user
     *
     * @tags Auth
     * @name LogoutCreate
     * @summary Logout user
     * @request POST:/auth/logout
     */
    logoutCreate: (Request: RouteLogoutBody, params: RequestParams = {}) =>
      this.request<any, any>({
        path: `/auth/logout`,
        method: "POST",
        body: Request,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Register User, Email has to be Unique and valid, Pseudo has to be Unique and > 3 characters, Password has to be > 8 characters, Name and surname has to be > 2 characters
     *
     * @tags Auth
     * @name RegisterCreate
     * @summary Register User
     * @request POST:/auth/register
     */
    registerCreate: (Request: RouteRegisterBody, params: RequestParams = {}) =>
      this.request<RouteRegisterReturn, any>({
        path: `/auth/register`,
        method: "POST",
        body: Request,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Renew Token via refresh token
     *
     * @tags Auth
     * @name RenewCreate
     * @summary Renew Token
     * @request POST:/auth/renew
     */
    renewCreate: (Request: RouteRenewTokenBody, params: RequestParams = {}) =>
      this.request<RouteRenewTokenReturn, any>({
        path: `/auth/renew`,
        method: "POST",
        body: Request,
        type: ContentType.Json,
        ...params,
      }),
  };
  chan = {
    /**
     * @description Get user channels
     *
     * @tags Chan
     * @name ChanList
     * @summary Get user channels
     * @request GET:/chan
     * @secure
     */
    chanList: (params: RequestParams = {}) =>
      this.request<RouteGetUserChanReturn, any>({
        path: `/chan`,
        method: "GET",
        secure: true,
        ...params,
      }),

    /**
     * @description Create channel, Name has to be unique
     *
     * @tags Chan
     * @name ChanCreate
     * @summary Create channel
     * @request POST:/chan
     * @secure
     */
    chanCreate: (Request: RouteCreateChanBody, params: RequestParams = {}) =>
      this.request<RouteCreateChanReturn, RouteCreateChanReturn>({
        path: `/chan`,
        method: "POST",
        body: Request,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Check if channel name is available
     *
     * @tags Chan
     * @name NameCreate
     * @summary Check if channel name is available
     * @request POST:/chan/name
     * @secure
     */
    nameCreate: (Request: RouteCheckChanNameBody, params: RequestParams = {}) =>
      this.request<RouteCheckChanNameReturn, RouteCheckChanNameReturn>({
        path: `/chan/name`,
        method: "POST",
        body: Request,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Get Public channel
     *
     * @tags Chan
     * @name PublicList
     * @summary Get Public channel
     * @request GET:/chan/public
     * @secure
     */
    publicList: (
      chanId: string,
      query?: { limit?: number; search?: number; page?: number },
      params: RequestParams = {}
    ) =>
      this.request<ModelChannel[], RouteGetOneChanReturn>({
        path: `/chan/public`,
        method: "GET",
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Get One Channel by id
     *
     * @tags Chan
     * @name ChanDetail
     * @summary Get One Channel by id
     * @request GET:/chan/{chanId}
     * @secure
     */
    chanDetail: (chanId: string, params: RequestParams = {}) =>
      this.request<ModelChannel, RouteGetOneChanReturn>({
        path: `/chan/${chanId}`,
        method: "GET",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Update Channel
     *
     * @tags Chan
     * @name ChanCreate2
     * @summary Update Channel
     * @request POST:/chan/{chanId}
     * @originalName chanCreate
     * @duplicate
     * @secure
     */
    chanCreate2: (
      chanId: number,
      body: RouteUpdateChannelBody,
      params: RequestParams = {}
    ) =>
      this.request<RouteUpdateChannelReturn, RouteUpdateChannelReturn>({
        path: `/chan/${chanId}`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Delete channel
     *
     * @tags Chan
     * @name ChanDelete
     * @summary Delete channel
     * @request DELETE:/chan/{chanId}
     * @secure
     */
    chanDelete: (chanId: string, params: RequestParams = {}) =>
      this.request<RouteDeleteChanReturn, RouteDeleteChanReturn>({
        path: `/chan/${chanId}`,
        method: "DELETE",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Get One Channel by id
     *
     * @tags Chan
     * @name JoinCreate
     * @summary Get One Channel by id
     * @request POST:/chan/{chanId}/join
     * @secure
     */
    joinCreate: (chanId: string, params: RequestParams = {}) =>
      this.request<ModelChannel, RouteJoinChanReturn>({
        path: `/chan/${chanId}/join`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Get One Channel messages by id, if user not in chan can't see message and if user hasn't the read right can only see past 24 hour message
     *
     * @tags Message
     * @name MessageDetail
     * @summary Get One Channel messages by id
     * @request GET:/chan/{chanId}/message
     * @secure
     */
    messageDetail: (
      chanId: string,
      query?: { limit?: number; page?: number },
      params: RequestParams = {}
    ) =>
      this.request<ModelMessage[], RouteGetOneChanReturn>({
        path: `/chan/${chanId}/message`,
        method: "GET",
        query: query,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Send Message
     *
     * @tags Message
     * @name MessageCreate
     * @summary Send Message
     * @request POST:/chan/{chanId}/message
     * @secure
     */
    messageCreate: (
      chanId: number,
      body: RouteSendMessageContent,
      params: RequestParams = {}
    ) =>
      this.request<RouteGetOneChanReturn, RouteGetOneChanReturn>({
        path: `/chan/${chanId}/message`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Reset channel password
     *
     * @tags Chan
     * @name RenewCreate
     * @summary Reset channel password
     * @request POST:/chan/{chanId}/renew
     * @secure
     */
    renewCreate: (chanId: string, params: RequestParams = {}) =>
      this.request<RouteRenewChanPasswordReturn, any>({
        path: `/chan/${chanId}/renew`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description Update User Channel Right
     *
     * @tags Chan
     * @name RightCreate
     * @summary Update User Channel Right
     * @request POST:/chan/{chanId}/right
     * @secure
     */
    rightCreate: (
      chanId: number,
      body: RouteUpdateUserChanRightBody,
      params: RequestParams = {}
    ) =>
      this.request<
        RouteUpdateUserChanRightReturn,
        RouteUpdateUserChanRightReturn
      >({
        path: `/chan/${chanId}/right`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),
  };
  user = {
    /**
     * @description Get user
     *
     * @tags User
     * @name UserList
     * @summary Get user
     * @request GET:/user
     * @secure
     */
    userList: (params: RequestParams = {}) =>
      this.request<ModelUser, any>({
        path: `/user`,
        method: "GET",
        secure: true,
        ...params,
      }),

    /**
     * @description Set user
     *
     * @tags User
     * @name UserCreate
     * @summary Set user
     * @request POST:/user
     * @secure
     */
    userCreate: (Request: RouteSetUserBody, params: RequestParams = {}) =>
      this.request<RouteSetUserReturn, any>({
        path: `/user`,
        method: "POST",
        body: Request,
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description SetPicture user
     *
     * @tags User
     * @name SetpictureCreate
     * @summary SetPicture user
     * @request POST:/user/setpicture
     * @secure
     */
    setpictureCreate: (data: { file: File }, params: RequestParams = {}) =>
      this.request<RouteSetPictureReturn, any>({
        path: `/user/setpicture`,
        method: "POST",
        body: data,
        secure: true,
        type: ContentType.FormData,
        ...params,
      }),
  };
}
