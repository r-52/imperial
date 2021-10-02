import { Component, EventEmitter, Output } from "@angular/core";
import { BaseComponent } from "../../../base/base/base.component";

@Component({
  selector: "app-create-button",
  templateUrl: "./create-button.component.html",
  styleUrls: ["./create-button.component.scss"],
})
export class CreateButtonComponent extends BaseComponent {
  @Output()
  public createClick: EventEmitter<void> = new EventEmitter<void>();

  public onCreateClick(): void {
    this.createClick.emit();
  }
}
