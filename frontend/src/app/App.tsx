import { HostsPanel } from "../features/hosts/HostsPanel";
import { OverviewPanel } from "../features/overview/OverviewPanel";
import { SessionsPanel } from "../features/sessions/SessionsPanel";
import { TrafficQualityPanel } from "../features/traffic-quality/TrafficQualityPanel";

export function App() {
  return (
    <main className="page">
      <header>
        <h1>Nginx Traffic Intelligence</h1>
        <p>Local-first dashboard skeleton for MVP implementation.</p>
      </header>
      <section className="grid">
        <OverviewPanel />
        <TrafficQualityPanel />
        <SessionsPanel />
        <HostsPanel />
      </section>
    </main>
  );
}
