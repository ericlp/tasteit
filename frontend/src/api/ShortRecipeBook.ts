import { Owner } from "./Owner";

export interface ShortRecipeBook {
  id: string;
  name: string;
  uniqueName: string;
  author: string;
  imageLink: string;
  uploadedBy: Owner;
}
