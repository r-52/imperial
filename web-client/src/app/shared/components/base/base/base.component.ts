import { Component, Injector, OnDestroy } from "@angular/core";
import { Subscription } from "rxjs";
import { DetailComponentModeEnum } from "../../../enums/detail-component/detail-component-mode.enum";
import { Title } from "@angular/platform-browser";

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

  constructor() {}

  ngOnDestroy() {
    for (const subscription of this._subscriptions) {
      try {
        subscription.unsubscribe();
      } catch (e) {}
    }
  }
}
