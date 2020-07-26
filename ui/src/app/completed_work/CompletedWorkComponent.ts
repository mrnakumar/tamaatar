import {Component, OnInit} from '@angular/core';
import {ApiService} from '../http/ApiService';
import {TimeBySprintName} from '../model/TimeBySprintName';

@Component({
  selector: 'app-completed-work',
  templateUrl: 'completed-work.component.html',
})
export class CompletedWorkComponent implements OnInit {
  completed: TimeBySprintName[];

  constructor(private api: ApiService) {
  }

  ngOnInit(): void {
    this.api.listSprints(sprints => {
      this.completed = sprints;
    });
  }
}
