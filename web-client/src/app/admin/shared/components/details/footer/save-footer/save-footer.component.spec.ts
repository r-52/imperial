import { ComponentFixture, TestBed } from '@angular/core/testing';

import { SaveFooterComponent } from './save-footer.component';

describe('SaveFooterComponent', () => {
  let component: SaveFooterComponent;
  let fixture: ComponentFixture<SaveFooterComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ SaveFooterComponent ]
    })
    .compileComponents();
  });

  beforeEach(() => {
    fixture = TestBed.createComponent(SaveFooterComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
