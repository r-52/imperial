import { Injectable } from "@angular/core";
import { environment } from "../../../../environments/environment";
import { Base } from "../../../admin/shared/models/api/base.model";

@Injectable({
  providedIn: "root",
})
export abstract class BaseService<T> {
  get baseUrl() {
    return environment.apiEndpoint;
  }

  public abstract get url(): string;

  public async getById(uid: string) {}
}
