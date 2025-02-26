export function escapeRegExp(string) {
  return string ? string.replace(/[.*+?^${}()|[\]\\]/g, "\\$&") : string;
}
