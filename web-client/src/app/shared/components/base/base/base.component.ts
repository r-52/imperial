import { Component, OnDestroy } from "@angular/core";
import { Subscription } from "rxjs";
import { DetailComponentModeEnum } from "../../../enums/detail-component/detail-component-mode.enum";

@Component({
  selector: "app-base",
  templateUrl: "./base.component.html",
  styleUrls: ["./base.component.scss"],
})
export class BaseComponent implements OnDestroy {
  /**
   * all the passed subscriptions for auto-unsub onDestroy
   * @protected
   */
  protected _subscriptions: Subscription[] = [];

  /**
   * the current mode that is used
   */
  public detailMode: DetailComponentModeEnum = DetailComponentModeEnum.newMode;

  /**
   * the type alias for the enum
   */
  public detailModeEnum = DetailComponentModeEnum;

  constructor() {}

  ngOnDestroy() {
    for (const subscription of this._subscriptions) {
      try {
        subscription.unsubscribe();
      } catch (e) {}
    }
  }
}
