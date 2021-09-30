import { ComponentFixture, TestBed } from "@angular/core/testing";

import { FabNewButtonComponent } from "./fab-new-button.component";

describe("FabNewButtonComponent", () => {
  let component: FabNewButtonComponent;
  let fixture: ComponentFixture<FabNewButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ FabNewButtonComponent ]
    })
      .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(FabNewButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("should create", () => {
    expect(component).toBeTruthy();
  });
});
