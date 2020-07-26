import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {catchError, retry} from 'rxjs/operators';

@Injectable()
export class ApiService {
  constructor(private http: HttpClient) {
  }

  public checkLogin(loggedIn: (next: boolean) => void): void {
    this.http.get('clogin').subscribe(x => {
      console.log(x);
      loggedIn(true);
    }, error => {
      loggedIn(false);
    });
  }

  public logIn(loggedIn: (next: boolean) => void): void {
    this.http.post('login', {
      Name: 'oho',
      Passwd: 'yahu',
    }).subscribe(next => {
      loggedIn(true);
    }, error => {
      loggedIn(false);
    });
  }
}
