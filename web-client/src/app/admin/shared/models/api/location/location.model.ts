import { Base } from "../base.model";

export interface Location extends Base {
  name: string;

  street1: string;

  street2: string;

  city: string;

  zip: string;

  country: string;
}
