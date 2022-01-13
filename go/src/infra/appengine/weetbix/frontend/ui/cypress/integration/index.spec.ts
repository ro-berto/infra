// Copyright 2021 The Chromium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style license that can be
// found in the LICENSE file.

describe('Index Page', () => {
    beforeEach(() => {
        cy.visit('/').get('button').click();
        cy.get('body').contains('Logout');
    })
    it('loads bugs table', () => {
        // Navigate to the bug cluster page
        cy.contains('Bugs').click();
        // check for the header text in the bug cluster table.
        cy.get('bugs-table').contains('Source Cluster ID');
    })
    it('loads cluster table', () => {
        // check for the header text in the cluster table.
        cy.get('cluster-table').contains('Test Runs Failed');
    })
    it('loads a cluster page', () => {
        cy.get('cluster-table').get('[data-cy=cluster-link]').first().click();
        cy.get('body').contains('Breakdown');
        // Check that the analysis section is showing at least one group.
        cy.get('td.group');
    })
})
