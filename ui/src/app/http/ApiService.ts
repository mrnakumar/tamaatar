import {HttpClient} from '@angular/common/http';
import {Injectable} from '@angular/core';
import {TimeBySprintName} from '../poto/TimeBySprintName';

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
      if (error.status === 401) {
        this.logIn(success => {
          if (success) {
            this.listSprints(sprintConsumer);
          } else {
            sprintConsumer([]);
          }
        });
      }
    });
  }

  public createSprint(name: string, duration: number, onSuccess: () => void): void {
    this.http.post('/createSprint', {
      Name: name,
      Duration: duration
    }).subscribe(_ => {
      onSuccess();
    }, error => {
      if (error.status === 401) {
        this.logIn(success => {
          if (success) {
            this.createSprint(name, duration, onSuccess);
          }
        });
      }
    });
  }
}
