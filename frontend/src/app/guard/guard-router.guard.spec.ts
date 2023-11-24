import { TestBed } from '@angular/core/testing';
import { CanActivateFn } from '@angular/router';

import { guardRouterGuard } from './guard-router.guard';

describe('guardRouterGuard', () => {
  const executeGuard: CanActivateFn = (...guardParameters) => 
      TestBed.runInInjectionContext(() => guardRouterGuard(...guardParameters));

  beforeEach(() => {
    TestBed.configureTestingModule({});
  });

  it('should be created', () => {
    expect(executeGuard).toBeTruthy();
  });
});
