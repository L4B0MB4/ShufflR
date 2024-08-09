export const fetchUtil = async <T>(
  path: string,
  method: string,
  body: any
): Promise<T | null> => {
  const res = await fetch(path, {
    method,
    body: body,
    credentials: "include",
  });

  if (res.status < 200 || res.status >= 400) {
    return null;
  }
  const resJson = await res.json();
  return resJson as T;
};
