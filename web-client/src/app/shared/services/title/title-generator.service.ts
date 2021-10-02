import { Injectable } from "@angular/core";
import { DetailComponentModeEnum } from "../../enums/detail-component/detail-component-mode.enum";

@Injectable({
  providedIn: "root",
})
export class TitleGeneratorService {
  constructor() {}

  public makeTitle(
    mode: DetailComponentModeEnum,
    name: string,
    module: string = ""
  ): string {
    if (mode === DetailComponentModeEnum.newMode) {
      return "New " + module;
    } else if (mode === DetailComponentModeEnum.editingMode) {
      return "Editing " + name;
    }
    return `${module} -${name}`;
  }
}
