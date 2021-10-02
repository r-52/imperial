import { Component, EventEmitter, Output } from "@angular/core";

@Component({
  selector: "app-cancel-button",
  templateUrl: "./cancel-button.component.html",
  styleUrls: ["./cancel-button.component.scss"],
})
export class CancelButtonComponent {
  @Output()
  public clickCancel: EventEmitter<void> = new EventEmitter<void>();

  constructor() {}

  public onClickCancel(): void {
    this.clickCancel.emit();
  }
}
