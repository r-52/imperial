import { Component, OnInit } from "@angular/core";
import { BaseComponent } from "../base/base.component";
import { DetailComponentModeEnum } from "../../../enums/detail-component/detail-component-mode.enum";

@Component({
  selector: "app-base-detail",
  templateUrl: "./base-detail.component.html",
  styleUrls: ["./base-detail.component.scss"],
})
export class BaseDetailComponent extends BaseComponent {
  /**
   * the title for the detail component
   */
  public title: string = "";

  /**
   * the current mode that is used
   */
  public detailMode: DetailComponentModeEnum = DetailComponentModeEnum.newMode;

  /**
   * the type alias for the enum
   */
  public detailModeEnum = DetailComponentModeEnum;

  constructor() {
    super();
  }

  ngOnInit(): void {}
}
