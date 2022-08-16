import { Image } from "./Image";
import { Owner } from "./Owner";

export interface RecipeBook {
  id: string;
  name: string;
  uniqueName: string;
  image: Image;
  recipes: RecipeBookRecipe[];
  uploadedBy: Owner;
  author: string;
}

export interface RecipeBookRecipe {
  id: string;
  name: string;
  uniqueName: string;
  author: string;
}
