import { Owner } from "../api/Owner";

export function isThisOwner(id: string, owners?: Owner[]) {
  return owners && owners.some((owner) => owner.id === id);
}
