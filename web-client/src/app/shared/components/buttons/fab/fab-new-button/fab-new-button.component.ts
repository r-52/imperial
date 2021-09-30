import { Component, EventEmitter, Output } from "@angular/core";

@Component({
  selector: "app-fab-new-button",
  templateUrl: "./fab-new-button.component.html",
  styleUrls: ["./fab-new-button.component.scss"],
})
export class FabNewButtonComponent {
  @Output()
  public newClick: EventEmitter<void> = new EventEmitter<void>();

  constructor() {}

  public onClick(): void {
    this.newClick.emit();
  }
}
