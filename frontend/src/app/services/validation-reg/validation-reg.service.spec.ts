import { TestBed } from '@angular/core/testing';

import { ValidationRegService } from './validation-reg.service';

describe('ValidationRegService', () => {
  let service: ValidationRegService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(ValidationRegService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
