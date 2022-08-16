import { RGBColor } from "./Color";
import { Owner } from "./Owner";

export interface Tag {
  id: string;
  name: string;
  description: string;
  color: RGBColor;
  recipeCount: number;
  author: Owner;
}
