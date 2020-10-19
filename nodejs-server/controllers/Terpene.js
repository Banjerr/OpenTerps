'use strict';

var utils = require('../utils/writer.js');
var Terpene = require('../service/TerpeneService');

module.exports.addTerpene = function addTerpene (req, res, next, body) {
  Terpene.addTerpene(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.deleteTerpene = function deleteTerpene (req, res, next, api_key, terpeneId) {
  Terpene.deleteTerpene(api_key, terpeneId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.findTerpenesByEffect = function findTerpenesByEffect (req, res, next, effect) {
  Terpene.findTerpenesByEffect(effect)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.getTerpeneById = function getTerpeneById (req, res, next, terpeneId) {
  Terpene.getTerpeneById(terpeneId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.updateTerpene = function updateTerpene (req, res, next, body) {
  Terpene.updateTerpene(body)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};

module.exports.updateTerpeneByID = function updateTerpeneByID (req, res, next, terpeneId) {
  Terpene.updateTerpeneByID(terpeneId)
    .then(function (response) {
      utils.writeJson(res, response);
    })
    .catch(function (response) {
      utils.writeJson(res, response);
    });
};
