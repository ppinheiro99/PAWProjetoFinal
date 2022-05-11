import { TestBed } from '@angular/core/testing';

import { PresentationsService } from './presentations.service';

describe('AuthService', () => {
  let service: PresentationsService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(PresentationsService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
