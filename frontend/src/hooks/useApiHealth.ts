import { useEffect, useState } from "react";
import { apiClient } from "../api/client";

type HealthState = "loading" | "ok" | "degraded";

export function useApiHealth() {
  const [state, setState] = useState<HealthState>("loading");

  useEffect(() => {
    let mounted = true;
    apiClient
      .health()
      .then((status) => {
        if (!mounted) return;
        setState(status === "ok" ? "ok" : "degraded");
      })
      .catch(() => {
        if (!mounted) return;
        setState("degraded");
      });

    return () => {
      mounted = false;
    };
  }, []);

  return state;
}
