import { Component, OnInit } from "@angular/core";
import { BaseComponent } from "../../../shared/components/base/base/base.component";
import { ActivatedRoute } from "@angular/router";
import { combineLatest, zip } from "rxjs";

@Component({
  selector: "app-location-details",
  templateUrl: "./location-details.component.html",
  styleUrls: ["./location-details.component.scss"],
})
export class LocationDetailsComponent extends BaseComponent implements OnInit {
  constructor(private readonly _route: ActivatedRoute) {
    super();
  }

  public ngOnInit() {
    this._subscriptions.push(
      combineLatest([this._route.data]).subscribe((next) => {
        const routeData = next[0];
        console.log(routeData);
      })
    );
  }
}
