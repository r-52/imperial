import { Component, EventEmitter, OnInit, Output } from "@angular/core";
import { BaseComponent } from "../../../base/base/base.component";

@Component({
  selector: "app-edit-button",
  templateUrl: "./edit-button.component.html",
  styleUrls: ["./edit-button.component.scss"],
})
export class EditButtonComponent extends BaseComponent {
  @Output()
  public editClick: EventEmitter<void> = new EventEmitter<void>();

  public onEditClick(): void {
    this.editClick.emit();
  }
}
