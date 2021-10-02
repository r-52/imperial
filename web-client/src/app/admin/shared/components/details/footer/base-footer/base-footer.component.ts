import { Component, Input } from "@angular/core";
import { DetailComponentModeEnum } from "../../../../../../shared/enums/detail-component/detail-component-mode.enum";
import { BaseComponent } from "../../../../../../shared/components/base/base/base.component";

@Component({
  selector: "app-base-footer",
  templateUrl: "./base-footer.component.html",
  styleUrls: ["./base-footer.component.scss"],
})
export class BaseFooterComponent extends BaseComponent {
  public detailModeEnum = DetailComponentModeEnum;

  @Input()
  public detailMode: DetailComponentModeEnum =
    DetailComponentModeEnum.editingMode;
}
