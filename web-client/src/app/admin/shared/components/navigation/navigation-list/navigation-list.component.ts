import { Component, OnInit } from "@angular/core";
import { NavigationEntryInterface } from "../../../models/sidebar/navigation-entry.interface";

@Component({
  selector: "app-navigation-list",
  templateUrl: "./navigation-list.component.html",
  styleUrls: ["./navigation-list.component.scss"],
})
export class NavigationListComponent {
  public navLinks: NavigationEntryInterface[] = [
    {
      iconName: "home",
      url: "/admin/",
      label: "Home",
    },
    {
      iconName: "supervisor_account",
      url: "/admin/persons",
      label: "Co-Worker",
    },
    {
      iconName: "location_city",
      url: "/admin/locations",
      label: "Locations",
    },
    {
      iconName: "text_snippet",
      url: "/admin/applications",
      label: "Applications",
    },
  ];

  constructor() {}
}
