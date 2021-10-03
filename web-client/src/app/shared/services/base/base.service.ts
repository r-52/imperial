import { Injectable } from "@angular/core";
import { environment } from "../../../../environments/environment";
import { Observable } from "rxjs";

@Injectable({
  providedIn: "root",
})
export abstract class BaseService<T> {
  get baseUrl() {
    return environment.apiEndpoint;
  }

  public abstract get url(): string;

  public getById(uid: string) {}

  public getAll(page: number, pageSize: number = 25) {}

  public create(obj: T) {}

  public update(uid: string, obj: T) {}
}
