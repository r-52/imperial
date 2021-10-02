import { Injectable } from "@angular/core";
import { BaseService } from "../base/base.service";
import { HttpClient } from "@angular/common/http";
import { Location as CompanyLocation } from "../../../admin/shared/models/api/location/location.model";

@Injectable({
  providedIn: "root",
})
export class LocationService extends BaseService<CompanyLocation> {
  constructor(private readonly _httpClient: HttpClient) {
    super();
  }

  public get url() {
    return `${this.baseUrl}/location/`;
  }

  public getLocations() {}

  public getLocationById(uid: string) {}
}
