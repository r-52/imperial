import { Component, OnDestroy, OnInit } from "@angular/core";
import { Subscription } from "rxjs";
import { ActivatedRoute } from "@angular/router";

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
