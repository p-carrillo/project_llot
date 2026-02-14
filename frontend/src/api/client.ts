type HealthResponse = {
  status: string;
};

const baseURL = import.meta.env.VITE_API_BASE_URL ?? "/api/v1";

export const apiClient = {
  async health(): Promise<string> {
    const response = await fetch(`${baseURL}/health`, {
      method: "GET",
      headers: {
        Accept: "application/json"
      }
    });

    if (!response.ok) {
      throw new Error(`Health endpoint failed with status ${response.status}`);
    }

    const body = (await response.json()) as HealthResponse;
    return body.status;
  }
};
