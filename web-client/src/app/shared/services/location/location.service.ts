import { Injectable } from "@angular/core";
import { BaseService } from "../base/base.service";

@Injectable({
  providedIn: "root",
})
export class LocationService extends BaseService {
  constructor() {
    super();
  }

  get url() {
    return `${this.baseUrl}/location/`;
  }
}
