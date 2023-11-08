/* tslint:disable */
/* eslint-disable */
/**
 * Ent Schema API
 * This is an auto generated API description made out of an Ent schema definition
 *
 * The version of the OpenAPI document: 0.1.0
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
/**
 * 
 * @export
 * @interface UserDevicesList
 */
export interface UserDevicesList {
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    id: string;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    name: string;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    description?: string;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    type: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserDevicesList
     */
    dns: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    publicKey: string;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    presharedKey: string;
    /**
     * 
     * @type {boolean}
     * @memberof UserDevicesList
     */
    keepAlive: boolean;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    endpoint: string;
    /**
     * 
     * @type {Array<string>}
     * @memberof UserDevicesList
     */
    allowedIps: Array<string>;
    /**
     * 
     * @type {string}
     * @memberof UserDevicesList
     */
    userId: string;
}

/**
 * Check if a given object implements the UserDevicesList interface.
 */
export function instanceOfUserDevicesList(value: object): boolean {
    let isInstance = true;
    isInstance = isInstance && "id" in value;
    isInstance = isInstance && "name" in value;
    isInstance = isInstance && "type" in value;
    isInstance = isInstance && "dns" in value;
    isInstance = isInstance && "publicKey" in value;
    isInstance = isInstance && "presharedKey" in value;
    isInstance = isInstance && "keepAlive" in value;
    isInstance = isInstance && "endpoint" in value;
    isInstance = isInstance && "allowedIps" in value;
    isInstance = isInstance && "userId" in value;

    return isInstance;
}

export function UserDevicesListFromJSON(json: any): UserDevicesList {
    return UserDevicesListFromJSONTyped(json, false);
}

export function UserDevicesListFromJSONTyped(json: any, ignoreDiscriminator: boolean): UserDevicesList {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'id': json['id'],
        'name': json['name'],
        'description': !exists(json, 'description') ? undefined : json['description'],
        'type': json['type'],
        'dns': json['dns'],
        'publicKey': json['public_key'],
        'presharedKey': json['preshared_key'],
        'keepAlive': json['keep_alive'],
        'endpoint': json['endpoint'],
        'allowedIps': json['allowed_ips'],
        'userId': json['user_id'],
    };
}

export function UserDevicesListToJSON(value?: UserDevicesList | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'id': value.id,
        'name': value.name,
        'description': value.description,
        'type': value.type,
        'dns': value.dns,
        'public_key': value.publicKey,
        'preshared_key': value.presharedKey,
        'keep_alive': value.keepAlive,
        'endpoint': value.endpoint,
        'allowed_ips': value.allowedIps,
        'user_id': value.userId,
    };
}
