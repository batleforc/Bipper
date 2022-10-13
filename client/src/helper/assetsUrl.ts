export const getAssetsUrl = (path: string) => {
  return `${import.meta.env.VITE_API}/asset/${path}`;
};
