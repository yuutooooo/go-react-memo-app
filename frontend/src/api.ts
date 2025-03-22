export const unauthenticatedApi = async (
  method: string,
  path: string,
  data?: any
) => {
  const response = await fetch(`${process.env.REACT_APP_API_URL}/${path}`, {
    method: method,
    body: data ? JSON.stringify(data) : undefined,
    headers: {
      "Content-Type": "application/json",
    },
    credentials: "include",
  });
  return response.json();
};

export const authenticatedApi = async (
  method: string,
  path: string,
  queryParams: string = "",
  data?: any
) => {
  let queryString = "";
  if (queryParams) {
    queryString = `?${queryParams}`;
  }

  const headers: Record<string, string> = {
    "Content-Type": "application/json",
  };

  const response = await fetch(
    `${process.env.REACT_APP_API_URL}/${path}${queryString}`,
    {
      method: method,
      body: data ? JSON.stringify(data) : undefined,
      headers: headers,
      credentials: "include",
    }
  );
  return response.json();
};
