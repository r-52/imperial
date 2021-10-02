import { TestBed } from '@angular/core/testing';

import { TitleGeneratorService } from './title-generator.service';

describe('TitleGeneratorService', () => {
  let service: TitleGeneratorService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(TitleGeneratorService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
