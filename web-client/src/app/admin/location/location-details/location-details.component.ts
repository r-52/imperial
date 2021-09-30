import { Component, OnInit } from "@angular/core";
import { BaseComponent } from "../../../shared/components/base/base/base.component";
import { ActivatedRoute } from "@angular/router";

@Component({
  selector: "app-location-details",
  templateUrl: "./location-details.component.html",
  styleUrls: ["./location-details.component.scss"],
})
export class LocationDetailsComponent extends BaseComponent implements OnInit {
  constructor() {
    super();
  }

  public ngOnInit() {
    super.ngOnInit();
  }
}
