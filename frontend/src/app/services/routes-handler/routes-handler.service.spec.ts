import { TestBed } from '@angular/core/testing';

import { RoutesHandlerService } from './routes-handler.service';

describe('RoutesHandlerService', () => {
  let service: RoutesHandlerService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(RoutesHandlerService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
