export const getQuerry = () => {
  let urlParams = new URLSearchParams(window.location.search);
  return urlParams;
};

export const getQuerryParam = (param: string) => {
  return getQuerry().get(param);
};

export const hasQuerryParam = (param: string) => {
  return getQuerry().has(param);
};
