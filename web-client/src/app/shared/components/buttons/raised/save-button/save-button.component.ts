import { Component, EventEmitter, OnInit, Output } from "@angular/core";
import { BaseComponent } from "../../../base/base/base.component";

@Component({
  selector: "app-save-button",
  templateUrl: "./save-button.component.html",
  styleUrls: ["./save-button.component.scss"],
})
export class SaveButtonComponent extends BaseComponent {
  @Output()
  public saveClick: EventEmitter<void> = new EventEmitter<void>();

  public onSaveClick() {
    this.saveClick.emit();
  }
}
