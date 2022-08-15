import { Owner } from "./Owner";
import { Tag } from "./Tag";

export interface ShortRecipe {
  id: string;
  name: string;
  uniqueName: string;
  imageLink: string;
  author: Owner;
  tags: Tag[];
  estimatedTime: number;
  numberOfIngredients: number;
}
