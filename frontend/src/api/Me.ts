import { Owner } from "./Owner";
import { User } from "./User";

export interface Me {
  user: User;
  owners: Owner[];
}
