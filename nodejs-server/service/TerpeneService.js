'use strict';


/**
 * Add a new terpene to the api
 * Add new terpene to the api
 *
 * body Terpene Terpene object
 * no response value expected for this operation
 **/
exports.addTerpene = function(body) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Deletes a terpene
 * Deletes a terpene from the API
 *
 * api_key String 
 * terpeneId Long Terpene id to delete
 * no response value expected for this operation
 **/
exports.deleteTerpene = function(api_key,terpeneId) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Finds Terpenes by effect
 * Multiple status values can be provided with comma separated strings
 *
 * effect List Effect values that need to be considered for filter
 * returns List
 **/
exports.findTerpenesByEffect = function(effect) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = [ {
  "tastes" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "smells" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "name" : "myrcene",
  "id" : 0,
  "category" : {
    "name" : "name",
    "id" : 6
  },
  "strains" : [ {
    "name" : "name",
    "id" : 2
  }, {
    "name" : "name",
    "id" : 2
  } ],
  "tags" : [ {
    "name" : "name",
    "id" : 1
  }, {
    "name" : "name",
    "id" : 1
  } ]
}, {
  "tastes" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "smells" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "name" : "myrcene",
  "id" : 0,
  "category" : {
    "name" : "name",
    "id" : 6
  },
  "strains" : [ {
    "name" : "name",
    "id" : 2
  }, {
    "name" : "name",
    "id" : 2
  } ],
  "tags" : [ {
    "name" : "name",
    "id" : 1
  }, {
    "name" : "name",
    "id" : 1
  } ]
} ];
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Find terpene by ID
 * Returns a single terpene
 *
 * terpeneId Long ID of terpene to return
 * returns Terpene
 **/
exports.getTerpeneById = function(terpeneId) {
  return new Promise(function(resolve, reject) {
    var examples = {};
    examples['application/json'] = {
  "tastes" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "smells" : [ {
    "name" : "name",
    "id" : 5
  }, {
    "name" : "name",
    "id" : 5
  } ],
  "name" : "myrcene",
  "id" : 0,
  "category" : {
    "name" : "name",
    "id" : 6
  },
  "strains" : [ {
    "name" : "name",
    "id" : 2
  }, {
    "name" : "name",
    "id" : 2
  } ],
  "tags" : [ {
    "name" : "name",
    "id" : 1
  }, {
    "name" : "name",
    "id" : 1
  } ]
};
    if (Object.keys(examples).length > 0) {
      resolve(examples[Object.keys(examples)[0]]);
    } else {
      resolve();
    }
  });
}


/**
 * Update an existing terpene
 * Update terpene in the api
 *
 * body Terpene Terpene object
 * no response value expected for this operation
 **/
exports.updateTerpene = function(body) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}


/**
 * Updates a terpene in the api
 * Update a single terpene in the API
 *
 * terpeneId Long ID of terpene that needs to be updated
 * no response value expected for this operation
 **/
exports.updateTerpeneByID = function(terpeneId) {
  return new Promise(function(resolve, reject) {
    resolve();
  });
}

