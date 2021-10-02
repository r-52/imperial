import { Component, EventEmitter, OnInit, Output } from "@angular/core";
import { BaseFooterComponent } from "../base-footer/base-footer.component";

@Component({
  selector: "app-save-footer",
  templateUrl: "./save-footer.component.html",
  styleUrls: ["./save-footer.component.scss"],
})
export class SaveFooterComponent extends BaseFooterComponent {
  @Output()
  public saveClick: EventEmitter<void> = new EventEmitter<void>();

  @Output()
  public cancelClick: EventEmitter<void> = new EventEmitter<void>();

  public onCancelClick(): void {
    this.cancelClick.emit();
  }

  public onSaveClick(): void {
    this.saveClick.emit();
  }
}
