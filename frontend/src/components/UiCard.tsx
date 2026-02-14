import { PropsWithChildren } from "react";

type UiCardProps = PropsWithChildren<{
  title: string;
}>;

export function UiCard({ title, children }: UiCardProps) {
  return (
    <article className="card" aria-label={title}>
      <h2>{title}</h2>
      {children}
    </article>
  );
}
