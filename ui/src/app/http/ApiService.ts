import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {Observable, throwError} from 'rxjs';
import {catchError, retry} from 'rxjs/operators';
import {TimeBySprintName} from '../model/TimeBySprintName';

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

  public listSprints(sprintConsumer: (next: TimeBySprintName[]) => void): void {
    this.http.get<TimeBySprintName[]>('timeBySprintName').subscribe(next => {
      sprintConsumer(next);
    }, error => {
      console.log(error);
      sprintConsumer([]);
    });
  }
}
