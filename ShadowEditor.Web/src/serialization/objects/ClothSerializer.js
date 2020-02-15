import BaseSerializer from '../BaseSerializer';
import Object3DSerializer from '../core/Object3DSerializer';
import Cloth from '../../object/component/Cloth';

/**
 * ClothSerializer
 * @author tengge / https://github.com/tengge1
 */
function ClothSerializer() {
    BaseSerializer.call(this);
}

ClothSerializer.prototype = Object.create(BaseSerializer.prototype);
ClothSerializer.prototype.constructor = ClothSerializer;

ClothSerializer.prototype.toJSON = function (obj) {
    var json = Object3DSerializer.prototype.toJSON.call(this, obj);

    return json;
};

ClothSerializer.prototype.fromJSON = function (json, parent, camera) { // eslint-disable-line
    var cloth = new Cloth();

    Object3DSerializer.prototype.fromJSON.call(this, json, cloth);

    return cloth;
};

export default ClothSerializer;