// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import {assert} from 'chai';
import {MrCommentContent} from './mr-comment-content.js';


let element;

describe('mr-comment-content', () => {
  beforeEach(() => {
    element = document.createElement('mr-comment-content');
    document.body.appendChild(element);

    document.body.style.setProperty('--mr-toggled-font-family', 'Some-font');
  });

  afterEach(() => {
    document.body.removeChild(element);

    document.body.style.removeProperty('--mr-toggled-font-family');
  });

  it('initializes', () => {
    assert.instanceOf(element, MrCommentContent);
  });

  it('changes rendered font based on --mr-toggled-font-family', async () => {
    element.content = 'A comment';

    await element.updateComplete;

    const fontFamily = window.getComputedStyle(element).getPropertyValue(
        'font-family');

    assert.equal(fontFamily, 'Some-font');
  });

  it('does not render spurious spaces', async () => {
    element.content =
      'Some text before a go/link and more text before <b>some bold text</b>.';

    await element.updateComplete;

    const textContents = Array.from(element.shadowRoot.children).map(
        (child) => child.textContent);

    assert.deepEqual(textContents, [
      'Some text before a',
      ' ',
      'go/link',
      ' and more text before ',
      'some bold text',
      '.',
    ]);

    assert.deepEqual(
        element.shadowRoot.textContent,
        'Some text before a go/link and more text before some bold text.');
  });

  it('does render markdown', async () => {
    element.prefs = new Map([['render_markdown', true]]);
    element.content = '### this is a header';
    element.projectName = 'monkeyrail';

    await element.updateComplete;

    const headerText = element.shadowRoot.querySelector('h3').textContent;
    assert.equal(headerText, 'this is a header');
  });

  it('does not render markdown when prefs are set to false', async () => {
    element.prefs = new Map([['render_markdown', false]]);
    element.projectName = 'monkeyrail';
    element.content = '### this is a header';

    await element.updateComplete;

    const commentText = element.shadowRoot.textContent;
    assert.equal(commentText, '### this is a header');
  });
});
