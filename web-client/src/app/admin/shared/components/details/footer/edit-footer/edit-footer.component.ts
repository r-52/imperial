import { Component, EventEmitter, OnInit, Output } from "@angular/core";
import { BaseComponent } from "../../../../../../shared/components/base/base/base.component";
import { BaseFooterComponent } from "../base-footer/base-footer.component";

@Component({
  selector: "app-edit-footer",
  templateUrl: "./edit-footer.component.html",
  styleUrls: ["./edit-footer.component.scss"],
})
export class EditFooterComponent extends BaseFooterComponent {
  @Output()
  public editClick: EventEmitter<void> = new EventEmitter<void>();

  @Output()
  public cancelClick: EventEmitter<void> = new EventEmitter<void>();

  onCancelClick(): void {
    this.cancelClick.emit();
  }

  onEditClick(): void {
    this.editClick.emit();
  }
}
