import {Injectable} from '@angular/core';

@Injectable()
export class LocalTagsService {
  private readonly localTagsKey = 'local_tags';

  updateOrGet(newTag: string) {
    const tagSet = new Set<string>();
    const tags = localStorage.getItem(this.localTagsKey);
    if (tags.trim().length > 0) {
      const tagsJson = JSON.parse(tags);
      for (const tag of tagsJson) {
        tagSet.add(tag);
      }
    }
    if (newTag) {
      tagSet.add(newTag.toLowerCase());
      localStorage.setItem(this.localTagsKey, JSON.stringify(Array.from(tagSet)));
    }
    return Array.from(tagSet).sort();
  }
}
