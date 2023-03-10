// Copyright 2019 The Chromium Authors
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

import {assert, expect} from 'chai';
import {MrAttachment} from './mr-attachment.js';
import {prpcClient} from 'prpc-client-instance.js';
import {FILE_DOWNLOAD_WARNING} from 'shared/settings.js';

let element;

describe('mr-attachment', () => {
  beforeEach(() => {
    element = document.createElement('mr-attachment');
    document.body.appendChild(element);
    sinon.stub(prpcClient, 'call').returns(Promise.resolve({}));
  });

  afterEach(() => {
    document.body.removeChild(element);
    prpcClient.call.restore();
  });

  it('initializes', () => {
    assert.instanceOf(element, MrAttachment);
  });

  it('shows image thumbnail', async () => {
    element.attachment = {
      thumbnailUrl: 'thumbnail.jpeg',
      contentType: 'image/jpeg',
    };
    await element.updateComplete;
    const img = element.shadowRoot.querySelector('img');
    assert.isNotNull(img);
    assert.isTrue(img.src.endsWith('thumbnail.jpeg'));
  });

  it('shows video thumbnail', async () => {
    element.attachment = {
      viewUrl: 'video.mp4',
      contentType: 'video/mpeg',
    };
    await element.updateComplete;
    const video = element.shadowRoot.querySelector('video');
    assert.isNotNull(video);
    assert.isTrue(video.src.endsWith('video.mp4'));
  });

  it('does not show image thumbnail if deleted', async () => {
    element.attachment = {
      thumbnailUrl: 'thumbnail.jpeg',
      contentType: 'image/jpeg',
      isDeleted: true,
    };
    await element.updateComplete;
    const img = element.shadowRoot.querySelector('img');
    assert.isNull(img);
  });

  it('does not show video thumbnail if deleted', async () => {
    element.attachment = {
      viewUrl: 'video.mp4',
      contentType: 'video/mpeg',
      isDeleted: true,
    };
    await element.updateComplete;
    const video = element.shadowRoot.querySelector('video');
    assert.isNull(video);
  });

  it('deletes attachment', async () => {
    prpcClient.call.callsFake(() => Promise.resolve({}));

    element.attachment = {
      attachmentId: 67890,
      isDeleted: false,
    };
    element.canDelete = true;
    element.projectName = 'proj';
    element.localId = 1234;
    element.sequenceNum = 3;
    await element.updateComplete;

    const deleteButton = element.shadowRoot.querySelector('chops-button');
    deleteButton.click();

    assert.deepEqual(prpcClient.call.getCall(0).args, [
      'monorail.Issues', 'DeleteAttachment',
      {
        issueRef: {
          projectName: 'proj',
          localId: 1234,
        },
        sequenceNum: 3,
        attachmentId: 67890,
        delete: true,
      },
    ]);
    assert.isTrue(prpcClient.call.calledOnce);
  });

  it('undeletes attachment', async () => {
    prpcClient.call.callsFake(() => Promise.resolve({}));
    element.attachment = {
      attachmentId: 67890,
      isDeleted: true,
    };
    element.canDelete = true;
    element.projectName = 'proj';
    element.localId = 1234;
    element.sequenceNum = 3;
    await element.updateComplete;

    const deleteButton = element.shadowRoot.querySelector('chops-button');
    deleteButton.click();

    assert.deepEqual(prpcClient.call.getCall(0).args, [
      'monorail.Issues', 'DeleteAttachment',
      {
        issueRef: {
          projectName: 'proj',
          localId: 1234,
        },
        sequenceNum: 3,
        attachmentId: 67890,
        delete: false,
      },
    ]);
    assert.isTrue(prpcClient.call.calledOnce);
  });

  it('view link is not displayed if not given', async () => {
    element.attachment = {};
    await element.updateComplete;
    const viewLink = element.shadowRoot.querySelector('.attachment-view');
    assert.isNull(viewLink);
  });

  it('view link is displayed if given', async () => {
    element.attachment = {
      viewUrl: 'http://example.com/attachment.foo',
    };
    await element.updateComplete;
    const viewLink = element.shadowRoot.querySelector('.attachment-view');
    assert.isNotNull(viewLink);
    expect(viewLink).to.be.displayed;
    assert.equal(viewLink.href, 'http://example.com/attachment.foo');
  });

  describe('download', () => {
    let downloadLink;

    beforeEach(async () => {
      sinon.stub(window, 'confirm').returns(false);


      element.attachment = {};
      await element.updateComplete;
      downloadLink = element.shadowRoot.querySelector('.attachment-download');
      // Prevent Karma from opening up new tabs because of simulated link
      // clicks.
      downloadLink.removeAttribute('target');
    });

    afterEach(() => {
      window.confirm.restore();
    });

    it('download link is not displayed if not given', async () => {
      element.attachment = {};
      await element.updateComplete;
      assert.isTrue(downloadLink.hidden);
    });

    it('download link is displayed if given', async () => {
      element.attachment = {
        downloadUrl: 'http://example.com/attachment.foo',
      };
      await element.updateComplete;
      const downloadLink = element.shadowRoot.querySelector(
          '.attachment-download');
      assert.isFalse(downloadLink.hidden);
      expect(downloadLink).to.be.displayed;
      assert.equal(downloadLink.href, 'http://example.com/attachment.foo');
    });

    it('download allows recognized file extension and type', async () => {
      element.attachment = {
        contentType: 'image/png',
        filename: 'not-a-virus.png',
        downloadUrl: '#',
      };
      await element.updateComplete;

      downloadLink.click();

      sinon.assert.notCalled(window.confirm);
    });

    it('file extension matching is case insensitive', async () => {
      element.attachment = {
        contentType: 'image/png',
        filename: 'not-a-virus.PNG',
        downloadUrl: '#',
      };
      await element.updateComplete;

      downloadLink.click();

      sinon.assert.notCalled(window.confirm);
    });

    it('download warns on unrecognized file extension and type', async () => {
      element.attachment = {
        contentType: 'application/virus',
        filename: 'fake-virus.exe',
        downloadUrl: '#',
      };
      await element.updateComplete;

      downloadLink.click();

      sinon.assert.calledOnce(window.confirm);
      sinon.assert.calledWith(window.confirm, FILE_DOWNLOAD_WARNING);
    });
  });
});
