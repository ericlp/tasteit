import { Image } from "./Image";
import { Ingredient } from "./Ingredient";
import { Owner } from "./Owner";
import { Step } from "./Step";

export interface EditRecipe {
  id: string;
  name: string;
  uniqueName: string;
  description: string;
  ovenTemperature: number;
  estimatedTime: number;
  steps: Step[];
  ingredients: Ingredient[];
  images: Image[];
  author: Owner;
  tags: string[];
  portions: number;
}
