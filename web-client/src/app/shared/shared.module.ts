import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { SpacerComponent } from "./components/layout/util/spacer/spacer.component";
import { FabNewButtonComponent } from "./components/buttons/fab/fab-new-button/fab-new-button.component";
import { MaterialModule } from "../material.module";
import { BaseComponent } from "./components/base/base/base.component";

@NgModule({
  declarations: [SpacerComponent, FabNewButtonComponent, BaseComponent],
  exports: [SpacerComponent, FabNewButtonComponent, BaseComponent],
  imports: [CommonModule, MaterialModule],
})
export class SharedModule {}
