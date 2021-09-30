import { Component } from "@angular/core";
import { Router } from "@angular/router";

@Component({
  selector: "app-location-list",
  templateUrl: "./location-list.component.html",
  styleUrls: ["./location-list.component.scss"],
})
export class LocationListComponent {
  constructor(private readonly _router: Router) {}

  public onNewClick(): void {
    this._router.navigateByUrl("/admin/location/new");
  }
}
