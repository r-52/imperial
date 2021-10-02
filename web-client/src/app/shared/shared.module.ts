import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { SpacerComponent } from "./components/layout/util/spacer/spacer.component";
import { FabNewButtonComponent } from "./components/buttons/fab/fab-new-button/fab-new-button.component";
import { MaterialModule } from "../material.module";
import { BaseComponent } from "./components/base/base/base.component";
import { CreateButtonComponent } from "./components/buttons/raised/create-button/create-button.component";
import { CancelButtonComponent } from "./components/buttons/basic/cancel-button/cancel-button.component";
import { EditButtonComponent } from "./components/buttons/raised/edit-button/edit-button.component";
import { SaveButtonComponent } from "./components/buttons/raised/save-button/save-button.component";

@NgModule({
  declarations: [
    SpacerComponent,
    FabNewButtonComponent,
    BaseComponent,
    CreateButtonComponent,
    CancelButtonComponent,
    EditButtonComponent,
    SaveButtonComponent,
  ],
  exports: [
    SpacerComponent,
    FabNewButtonComponent,
    BaseComponent,
    CreateButtonComponent,
    CancelButtonComponent,
    EditButtonComponent,
    SaveButtonComponent,
  ],
  imports: [CommonModule, MaterialModule],
})
export class SharedModule {}
