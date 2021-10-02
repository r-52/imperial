import { Component, EventEmitter, Input, Output } from "@angular/core";
import { DetailComponentModeEnum } from "../../../../../../shared/enums/detail-component/detail-component-mode.enum";
import { BaseFooterComponent } from "../base-footer/base-footer.component";

@Component({
  selector: "app-create-footer",
  templateUrl: "./create-footer.component.html",
  styleUrls: ["./create-footer.component.scss"],
})
export class CreateFooterComponent extends BaseFooterComponent {
  @Output()
  public createClick: EventEmitter<void> = new EventEmitter<void>();

  @Output()
  public cancelClick: EventEmitter<void> = new EventEmitter<void>();

  /* **** Event listener **** */

  public onCancelClick(): void {
    this.cancelClick.emit();
  }

  public onCreateClick(): void {
    this.createClick.emit();
  }
}
