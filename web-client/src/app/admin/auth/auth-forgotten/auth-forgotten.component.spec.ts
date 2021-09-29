import { ComponentFixture, TestBed } from "@angular/core/testing";

import { AuthForgottenComponent } from "./auth-forgotten.component";

describe("AuthForgottenComponent", () => {
  let component: AuthForgottenComponent;
  let fixture: ComponentFixture<AuthForgottenComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ AuthForgottenComponent ]
    })
      .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(AuthForgottenComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it("should create", () => {
    expect(component).toBeTruthy();
  });
});
